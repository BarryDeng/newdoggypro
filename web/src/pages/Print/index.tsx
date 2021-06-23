import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import PrintPanel from '@/components/Demo'
import DisplayPanel from '@/components/Demo/DisplayPanel';

const Print: React.FC =  () => {
    return (
        <PageContainer>
            <PrintPanel 
                color="#ffc600"
                width={1000}
                height={1000}
                brushRadius={10}
                lazyRadius={12}
            />
            <DisplayPanel 
                color="#ffc600"
                width={1000}
                height={1000}
                brushRadius={10}
                lazyRadius={12}
            />
        </PageContainer>
    )
}

export default Print;